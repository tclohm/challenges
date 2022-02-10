class Promise:
	def __init__(self):
		self._on_resolve = []
		self._resolved = False
		self._value = None

	def then(self, callback):
		if self._resolved:
			callback(*self._value)
		else:
			self._on_resolve.append(callback)

	def _resolve(self, *args):
		if not self._resolved:
			self._resolved = True
			self._value = args
			for callback in self._on_resolve:
				callback(*args)


class Promise(Context):
  def __init__(self):
    self._on_resolve = []
    self._on_reject = []
    self._resolved = False
    self._rejected = False
    self._value = None

  def then(self, cb):
    if self._resolved:
      self.evloop._execute(cb, *self._value)
    elif not self._rejected:
      self._on_resolve.append(cb)
    return self

  def catch(self, cb):
    if self._rejected:
      self.evloop._execute(cb, self._value)
    elif not self._resolved:
      self._on_reject.append(cb)
    return self

  def _resolve(self, *args):
    if self._resolved or self._rejected:
      return

    self._resolved = True
    self._value = args
    for cb in self._on_resolve:
      self.evloop._execute(cb, *args)

  def _reject(self, err):
    if self._resolved or self._rejected:
      return
    self._rejected = True
    self._value = err
    for cb in self._on_reject:
      cb(err)


# A replacement for set_timer()
def sleep(duration):
  return Context._event_loop.set_timer(duration * 10e3)


class socket(Context):
  def __init__(self, *args):
    self._sock = _socket.socket(*args)
    self._sock.setblocking(False)
    self.evloop.register_fileobj(self._sock, self._on_event)
    # 0 - initial
    # 1 - connecting
    # 2 - connected
    # 3 - closed
    self._state = 0
    self._callbacks = {}

  def connect(self, addr):
    assert self._state == 0
    self._state = 1

    p = Promise()
    def _on_conn(err):
      if err:
        p._reject(err)
      else:
        p._resolve()

    self._callbacks['conn'] = _on_conn
    err = self._sock.connect_ex(addr)
    assert errno.errorcode[err] == 'EINPROGRESS'
    return p

  def recv(self, n):
    assert self._state == 2
    assert 'recv' not in self._callbacks

    p = Promise()
    def _on_read_ready(err):
      if err:
        p._reject(err)
      else:
        data = self._sock.recv(n)
        p._resolve(data)

    self._callbacks['recv'] = _on_read_ready
    return p

  def sendall(self, data):
    assert self._state == 2
    assert 'sent' not in self._callbacks

    p = Promise()
    def _on_write_ready(err):
      nonlocal data
      if err:
        return p._reject(err)

      n = self._sock.send(data)
      if n < len(data):
        data = data[n:]
        self._callbacks['sent'] = _on_write_ready
      else:
        p._resolve(None)

    self._callbacks['sent'] = _on_write_ready
    return p

  # ...
import time

class EventLoop:
	def __init__(self):
		self._queue = Queue()
		self._time = None

	def run(self, entry_point, *args):
		self._execute(entry_point, *args)

		while not self._queue.is_empty():
			fn, mask = self._queue.pop(self._time)
			self._execute(fn, mask)

		self._queue.close()

	def _execute(self, callback, *args):
		self._time = hrtime()
		try:
			callback(*args)
		except Exception as err:
			print('Uncaught exception:', err)
		self._time = hrtime()

	def register_fileobj(self, fileobj, callback):
		self._queue.register_fileobj(fileobj, callback)

	def unregister_fileobj(self, fileobj):
		self._queue.unregister_fileobj(fileobj)

	def set_timer(self, duration, callback):
		self._time = hrtime()
		self._queue.register_timer(self._time + duration, lambda _: callback())

def hrtime():
	return int(time.time() * 10e6)
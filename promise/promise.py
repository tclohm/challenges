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
		
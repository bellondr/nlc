class MyCircularQueue:

    def __init__(self, k: int):
        self.cap = k
        self.data = []

    def enQueue(self, value: int) -> bool:
        if len(self.data) < self.cap:
            self.data.append(value)
            return True
        return False

    def deQueue(self) -> bool:
        if len(self.data) == 0:
            return False
        else:
            self.data = self.data[1:]
            return True

    def Front(self) -> int:
        if len(self.data) == 0:
            return -1
        return self.data[0]

    def Rear(self) -> int:
        if len(self.data) == 0:
            return -1
        else:
            return self.data[-1]

    def isEmpty(self) -> bool:
        return len(self.data) == 0

    def isFull(self) -> bool:
        return len(self.data) == self.cap

# Your MyCircularQueue object will be instantiated and called as such:
# obj = MyCircularQueue(k)
# param_1 = obj.enQueue(value)
# param_2 = obj.deQueue()
# param_3 = obj.Front()
# param_4 = obj.Rear()
# param_5 = obj.isEmpty()
# param_6 = obj.isFull()
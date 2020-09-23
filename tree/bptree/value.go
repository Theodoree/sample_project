package bptree

type value struct {
    val interface{}
}

func NewValue() *value {
    return &value{}
}

func (v *value) reset() {
    v.val = nil
}

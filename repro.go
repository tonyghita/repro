package repro

type A struct{}

func (*A) F() (_ *B) { return }

type B struct{}

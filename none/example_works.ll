define void @f() {
  call i32 @a()
  add i32 0, 0
  call double @a()
  ret void
}

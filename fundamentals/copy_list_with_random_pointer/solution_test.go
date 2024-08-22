package copy_list_with_random_pointer

import "testing"

func TestExample1(t *testing.T) {
	head1 := &Node{
		Val: 7,
		Next: &Node{
			Val: 13,
			Next: &Node{
				Val: 11,
				Next: &Node{
					Val: 10,
					Next: &Node{
						Val: 1,
					}}}}}
	head1.Random = nil
	head1.Next.Random = head1
	head1.Next.Next.Random = head1.Next.Next.Next.Next
	head1.Next.Next.Next.Random = head1.Next.Next
	head1.Next.Next.Next.Next.Random = head1

	head2 := CopyRandomList(head1)

	curr1 := head1
	curr2 := head2
	for curr1 != nil {
		if curr1 == curr2 {
			t.Fail()
		}
		if curr1.Val != curr2.Val {
			t.Fail()
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
	if head2.Random != nil {
		t.Fail()
	}
	if head2.Next.Random != head2 {
		t.Fail()
	}
	if head2.Next.Next.Random != head2.Next.Next.Next.Next {
		t.Fail()
	}
	if head2.Next.Next.Next.Random != head2.Next.Next {
		t.Fail()
	}
	if head2.Next.Next.Next.Next.Random != head2 {
		t.Fail()
	}
}

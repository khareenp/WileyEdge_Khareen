package assignments

import "fmt"
type Node struct{
    Value int
}

type Stack struct{
    nodes[] *Node
}
type Queue struct{
    nodes[] *Node
}
func NewStack() *Stack{
    return &Stack{}
}

func NewQueue() *Queue{
    return &Queue{}
}

// Push : Adds an element on the stack
func (s *Stack)  Push(n *Node){
         s.nodes=append(s.nodes,n)
         
    }

// Pop : removes an element from the stack and returns its value
func (s *Stack)  Pop()  *Node{
         if len(s.nodes)==0{ 
            return nil
        } 
        n:= s.nodes[len(s.nodes)-1]
        s.nodes= s.nodes[:len(s.nodes)-1]
        return n
    } 

func (q *Queue)  Enqueue(n *Node){
        q.nodes=append(q.nodes,n)
    }

func (q *Queue)  Dequeue()  *Node{
        if len(q.nodes)==0{ 
        return nil
    }
        n:= q.nodes[0]
		 q.nodes= q.nodes[1:]
         //q.count++
        return  n
    }

func (n  *Node) String()string{
    return fmt.Sprint(n.Value)
}
func SnQ(){
   s:=NewStack()
   q:=NewQueue()
   
   s.Push( &Node{10} )
   s.Push( &Node{12} )
   s.Push( &Node{14} )
   s.Push( &Node{16} )

   q.Enqueue( &Node{20} )
   q.Enqueue( &Node{30} )
   q.Enqueue( &Node{40} )
   q.Enqueue( &Node{50} )

    fmt.Println(s.Pop()  ,s.Pop(),s.Pop(),s.Pop())
    fmt.Println(q.Dequeue()  ,q.Dequeue(),q.Dequeue(),q.Dequeue())

}


import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import Link from 'next/link';
import { useState } from 'react';
import { Todo } from '../../type/todo';
import axios from "axios";
import { Button, Container, Form } from 'react-bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css';

const Update: NextPage = () => {
  const router = useRouter();
  const todoValue = String(router.query.todo)
  const todoId = String(router.query.id);
  const [todo, setTodo] = useState<Todo>({ id: todoId, todo: todoValue })

  const updateTodo = (e: any) => {
    const id = e.target.value
    setTodo(prevState => ({ ...prevState, id: id }))
    axios.put("http://localhost:8000/todos/update/" + id, todo).then(res => {
      router.push("/")
    }).catch(err => {
      console.log(err)
    })
  }

  const changeValue = (e: any) => {
    const formValue = e.target.value
    setTodo(prevState => ({ ...prevState, todo: formValue }))
  }

  return (
    <>
      <Container>
        <h2 style={{ marginTop: "10px" }}>Update</h2>
        <div>
          <Form.Control type="text" value={todo.todo} onChange={changeValue}
            style={{ maxWidth: "40%" }}>
          </Form.Control>
          <Link href="/">
            <Button variant="primary" style={{ marginTop: "10px", marginRight: "10px" }}>Back</Button>
          </Link>
          <Button variant="success" style={{ marginTop: "10px" }} onClick={updateTodo} value={todo.id}>更新</Button>
        </div>
      </Container>
    </>
  )
}

export default Update

import type { NextPage } from 'next'
import Link from 'next/link';
import React, { useEffect, useState } from 'react'
import { Todo } from '../type/todo';
import axios from "axios";
import { Button, Table, Form, Container, Modal } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

const Home: NextPage = () => {
  const [todos, setTodos] = useState<Todo[]>([])
  const [todo, setTodo] = useState<Todo>({ id: "", todo: "" })
  const [mounted, setMounted] = useState<boolean>(false)
  const [showModal, setShowModal] = useState<boolean>(false)

  useEffect(() => {
    // 初回に二度実行される問題
    (
      async () => {
        const data = await axios.get("http://localhost:8000/todos")
        setTodos(data.data)
      }
    )()
    setMounted(true);
  }, [])

  const addTodo = (e: any) => {
    axios.post("http://localhost:8000/todos", todo).then(res => {
      setTodos(res.data)
    }).catch(err => {
      console.log(err)
    })
  }

  const deleteTodo = (e: any) => {
    const id = e.target.value

    axios.delete("http://localhost:8000/todos/" + id,).then(res => {
      setTodos(res.data)
    }).catch(err => {
      console.log(err)
    })
    setShowModal(false);
  }

  const handleShow = (e: any) => {
    const id = e.target.value
    setTodo(prevState => ({ ...prevState, id: id }))
    setShowModal(true);
  }
  const handleClose = () => setShowModal(false);

  const changeTodo = (e: any) => {
    const todo = e.target.value
    setTodo(prevState => ({ ...prevState, todo: todo }))
  }

  return (
    <Container>
      {mounted ? (
        <>
          <div>
            <h2 style={{ marginTop: "10px" }}>Todoアプリ</h2>
            <Form.Control type='text' value={todo.todo} onChange={changeTodo}
              style={{ width: '25%', marginTop: "10px", marginRight: "10px", display: "inline" }} />
            <Button variant="primary" onClick={addTodo}>追加</Button>
          </div>
          <div className="justify-content-md-center">
            <Modal show={showModal} onHide={handleClose}>
              <Modal.Header closeButton>
                <Modal.Title>削除</Modal.Title>
              </Modal.Header>
              <Modal.Body>このタスクを削除しますか？</Modal.Body>
              <Modal.Footer>
                <Button variant="secondary" onClick={handleClose}>いいえ</Button>
                <Button variant="success" onClick={deleteTodo} value={todo.id}>はい</Button>
              </Modal.Footer>
            </Modal>
          </div>
          <Table striped hover style={{ maxWidth: "60%" }}>
            <thead>
              <tr>
                <th style={{ textAlign: "center" }}>Todo</th>
                <th style={{ textAlign: "center" }}>更新</th>
                <th style={{ textAlign: "center" }}>削除</th>
              </tr>
            </thead>
            <tbody>
              {todos.map(v => (
                <tr key={v.id}>
                  <td>{v.todo}</td>
                  <td >
                    <Link as={`/update/${v.id}`} href={{ pathname: `/update/[id]`, query: v }}>
                      <Button variant="success" style={{ display: "block", margin: "auto" }}>更新</Button>
                    </Link>
                  </td>
                  <td><Button variant="danger" onClick={handleShow} value={v.id}
                    style={{ display: "block", margin: "auto" }}>削除</Button>
                  </td>
                </tr>
              ))}
            </tbody>
          </Table>
        </>) : <div></div>
      }
    </Container >
  )
}

export default Home

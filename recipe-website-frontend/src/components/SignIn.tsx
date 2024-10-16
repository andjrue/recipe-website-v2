import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Container, Row, Col, Card, Button, Form } from 'react-bootstrap';
import axios from 'axios'
import { useNavigate } from 'react-router-dom';

const BASE_URL = 'http://localhost:6969'

function SignIn() {
    const [user, setUser] = useState({ email: "", username: "", password: ""})
    const navigate = useNavigate()

    const handleInputChange = (e) => {
        const { name, value } = e.target
        setUser({ ...user, [name]: value })

    }

    const handleSignIn = async () => {
        try {
            await axios.post(`${BASE_URL}/users`, user)
            setUser({ email: "", username: "", password: ""})
            navigate('/recipes');
    } catch (error) {
      console.error('Error signing in:', error);
    }
  };

  return (
    <Container>
      <Row className="mt-5">
        <Col>
          <h1>Sign In</h1>
          <Card className="mt-3">
            <Card.Body>
              <Form>
                <Form.Group controlId="formEmail">
                  <Form.Label>Email</Form.Label>
                  <Form.Control
                    type="email"
                    name="email"
                    value={user.email}
                    onChange={handleInputChange}
                    placeholder="Enter email"
                  />
                </Form.Group>

                <Form.Group controlId="formUsername" className="mt-3">
                  <Form.Label>Username</Form.Label>
                  <Form.Control
                    type="text"
                    name="username"
                    value={user.username}
                    onChange={handleInputChange}
                    placeholder="Enter username"
                  />
                </Form.Group>

                <Form.Group controlId="formPassword" className="mt-3">
                  <Form.Label>Password</Form.Label>
                  <Form.Control
                    type="password"
                    name="password"
                    value={user.password}
                    onChange={handleInputChange}
                    placeholder="Enter password"
                  />
                </Form.Group>

                <Button variant="primary" className="mt-3" onClick={handleSignIn}>
                  Sign In
                </Button>
              </Form>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
}

export default SignIn;


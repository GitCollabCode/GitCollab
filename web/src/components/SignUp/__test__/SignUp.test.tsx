import React from 'react'
import { render, screen } from '@testing-library/react'

import Login from '../../Login/Login'

const mockedUsedNavigate = jest.fn()
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedUsedNavigate,
}))

describe('SignUp', () => {
  it('Test Render', () => {
    render(<Login />)
    const signUpText = screen.getAllByText('Sign Up')
    expect(signUpText.length).toBe(2)
    expect(screen.getByText('First name')).toBeInTheDocument()
    expect(screen.getByText('Last name')).toBeInTheDocument()
    expect(screen.getByText('Email address')).toBeInTheDocument()
    expect(screen.getByText('Password')).toBeInTheDocument()
  })
})

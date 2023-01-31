import React from 'react'
import { fireEvent, render, screen, waitFor } from '@testing-library/react'

import Button from '../Button'

const mockedOnClick = jest.fn()

describe('Button', () => {
  it('Test Button with new type', () => {
    render(<Button type="new" text="Button text" />)
    expect(screen.getByText('Button text')).toBeInTheDocument()
    expect(screen.getByTitle('new')).toBeInTheDocument()
  })

  it('Test Button with login type', () => {
    render(<Button type="login" text="Button text" />)
    expect(screen.getByText('Button text')).toBeInTheDocument()
    expect(screen.getByTitle('login')).toBeInTheDocument()
  })

  it('Test Button with logout type', () => {
    render(<Button type="logout" text="Button text" />)
    expect(screen.getByText('Button text')).toBeInTheDocument()
    expect(screen.getByTitle('logout')).toBeInTheDocument()
  })

  it('Test Button with next type', () => {
    render(<Button type="next" text="Button text" />)
    expect(screen.getByText('Button text')).toBeInTheDocument()
    expect(screen.getByTitle('next')).toBeInTheDocument()
  })

  it('Test Button onClick', async () => {
    render(<Button type="next" text="Button text" onClick={mockedOnClick} />)
    const button = screen.getByText('Button text')

    fireEvent.click(button)
    await waitFor(() => expect(mockedOnClick).toHaveBeenCalled())
  })
})

import React from 'react'
import {  render, screen } from '@testing-library/react'

import Footer from '../Footer'


describe('Footer', () => {
  it('Test Footer renders', () => {
    render(<Footer/>)
    expect(screen.getByTestId("footer")).toBeInTheDocument()
  })
})

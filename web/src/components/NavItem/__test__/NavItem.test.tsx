import React from 'react'
import { render, screen } from '@testing-library/react'

import NavItem from '../NavItem'

describe('NavItem', () => {
  it('Test NavItem render', () => {
    render(<NavItem text="test text" link="urlLink" />)
    expect(screen.getByText('test text')).toBeInTheDocument()
    expect(screen.getByText('test text').closest('a')).toHaveAttribute(
      'href',
      'urlLink'
    )
  })
})

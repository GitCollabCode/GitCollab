import React from 'react'
import { render, screen } from '@testing-library/react'

import LoadingSpinner from '../LoadingSpinner'

describe('LoadingSpinner', () => {
  it('Test LoadingSpinner renders with full overlay', () => {
    render(<LoadingSpinner isLoading={true} />)
    expect(screen.getByTestId('loader')).toBeInTheDocument()
    expect(screen.getByTestId('overlay')).toBeInTheDocument()
  })

  it('Test LoadingSpinner renders with small overlay', () => {
    render(<LoadingSpinner isLoading={true} type={'fixed'} />)
    expect(screen.getByTestId('loader')).toBeInTheDocument()
    expect(screen.getByTestId('dynamicOverlay')).toBeInTheDocument()
  })
})

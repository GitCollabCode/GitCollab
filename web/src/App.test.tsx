import React from 'react'
import { fireEvent, render, screen, waitFor } from '@testing-library/react'
import App from './App'
import { act } from 'react-dom/test-utils'

const mockMatches = true

window.matchMedia = (query) => ({
  matches: mockMatches,
  media: query,
  onchange: null,
  addListener: jest.fn(), // deprecated
  removeListener: jest.fn(), // deprecated
  addEventListener: jest.fn(),
  removeEventListener: jest.fn(),
  dispatchEvent: jest.fn(),
})

test('renders', () => {
  render(<App />)
  act(() => {
    fireEvent.click(screen.getByText('Open Menu'))
  })

  waitFor(() => {
    expect(screen.getByText('discover')).toBeInTheDocument()
  })

  const links = document.getElementsByClassName('link')
  expect(links.length).toBe(3)
})

import React, { CSSProperties } from 'react'
import RingLoader from 'react-spinners/RingLoader'
import style from './LoadingSpinner.module.css'

const LoadingSpinner = ({
  isLoading,
  type,
}: {
  isLoading: boolean
  type?: 'fixed'
}) => {
  const override: CSSProperties = {
    position: 'fixed',
    margin: '0 auto',
    top: '50%',
    left: 'calc(50% - 75px)',
    alignSelf: 'center',
    justifySelf: 'center',
  }

  return (
    <div
      className={type !== 'fixed' ? style.overlay : style.dynamicOverlay}
      data-testid={type !== 'fixed' ? 'overlay' : 'dynamicOverlay'}
    >
      <RingLoader
        loading={isLoading}
        color={'#ffffff'}
        cssOverride={override}
        size={150}
        aria-label="Loading Spinner"
        data-testid="loader"
      />
    </div>
  )
}

export default LoadingSpinner

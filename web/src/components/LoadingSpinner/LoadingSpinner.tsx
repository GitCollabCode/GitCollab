import React, { CSSProperties } from 'react'
import RingLoader from 'react-spinners/RingLoader'
import style from './LoadingSpinner.module.css'

const LoadingSpinner = ({ isLoading }: { isLoading: boolean }) => {
  const override: CSSProperties = {
    position: 'fixed',
    margin: '0 auto',
    top: '50%',
    left: 'calc(50% - 75px)',
    alignSelf: 'center',
    justifySelf: 'center',
  }

  console.log(isLoading)
  return (
    <div className={style.overlay}>
      heelllo
      <RingLoader
        loading={true}
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

import React, {createRef, useEffect, useState} from 'react'
import axios from 'axios'
import {toast} from 'react-toastify'
import Paper from '@material-ui/core/Paper'
import cx from 'classnames'
import Grid from '@material-ui/core/Grid'
import TextField from '@material-ui/core/TextField'
import InputAdornment from '@material-ui/core/InputAdornment'
import Button from '@material-ui/core/Button'
import capitalize from '@material-ui/core/utils/capitalize'
import NumberFormat from "react-number-format"
import {getColorStyle, TEST_USER_ID} from '../App'
import {DUX_TRADE_URL} from '../urls'

const SelectedOffer = ({selectedOfferId, getSelectedOffer, balance, setUpdateBalance}) => {
  const isSelected = selectedOfferId !== undefined
  const selectedOffer = isSelected && getSelectedOffer(selectedOfferId)
  const {id, unitPrice, unitType: {color} = {}, volume} = selectedOffer
  const maxQuantity = balance/unitPrice
  const [volumeToBuy, setVolumeToBuy] = useState('')
  const inputRef = createRef()

  useEffect(() => {
    if (!isSelected) {
      setVolumeToBuy('')
    } else {
      inputRef.current.focus()
    }
  },[inputRef, isSelected])

  const handleBuyButtonOnClick = ({userId}) => {
    axios.post(DUX_TRADE_URL,  {
      userId: TEST_USER_ID,
      offerId: id,
      volume: parseInt(volumeToBuy)
    }).then((response) => {
      setUpdateBalance(true)
      toast.info('Order Filled!')
    }, (error) => {
      toast.error('Error, Order Cancelled!')
      console.error(error)
    })
  }

  return (
    <Paper className={cx('border-blue',{'offer-details-active': isSelected})}>
      <header className='header'>
        <h3>Offer Details</h3>
      </header>
      <Grid container spacing={1} className='w-100 selected-offer-grid'>
        <OfferDetails {...{id, color, unitPrice, volume}} />
        <OrderControls
          {...{volumeToBuy, isSelected, setVolumeToBuy, inputRef, maxQuantity, unitPrice, handleBuyButtonOnClick}}
        />
      </Grid>
    </Paper>
  )
}

const OfferDetails = ({id, color, unitPrice, volume}) => (
  <>
    <Grid item xs={4}>
      Offer ID:
    </Grid>
    <Grid item xs={8}>
      {id}
    </Grid>
    <Grid item xs={4}>
      Color:
    </Grid>
    <Grid item xs={8} style={getColorStyle(color)}>
      {color && capitalize(color)}
    </Grid>
    <Grid item xs={4}>
      Price:
    </Grid>
    <Grid item xs={8}>
      £ {unitPrice}
    </Grid>
    <Grid item xs={4}>
      Volume
    </Grid>
    <Grid item xs={8}>
      {volume}
    </Grid>
  </>
)

const OrderControls = (props) => {
  const {volumeToBuy, isSelected, setVolumeToBuy, inputRef, maxQuantity, unitPrice, handleBuyButtonOnClick} = props

  return (
    <Grid item xs={12} className='order-controls'>
      <Grid container spacing={2} alignItems='center' justify='space-between'>
        <Grid item xs={6} sm={4}>
          <TextField
            style={{textAlign: 'right'}}
            startadornment={<InputAdornment position="start">£</InputAdornment>}
            value={parseInt(volumeToBuy, 10)}
            disabled={!isSelected}
            onChange={event => setVolumeToBuy(event.target.value)}
            variant="outlined"
            label='Quantity'
            inputProps={{
              ref: inputRef,
              max: maxQuantity,
              style: {textAlign: 'right'}
            }}
            InputProps={{inputComponent: NumberFormatCustom}}
          />
        </Grid>
        {isSelected &&
          <Grid item xs={6} sm={5}>
            <OrderCalculation {...{volumeToBuy, unitPrice}} />
          </Grid>
        }
        <Grid item xs={12} sm={3}>
          <Button
            className='buy-button'
            color={'primary'}
            variant='contained'
            disabled={!isSelected || volumeToBuy === ''}
            onClick={handleBuyButtonOnClick}
          >
            Buy
          </Button>
        </Grid>
      </Grid>
    </Grid>
  )
}

const OrderCalculation = ({volumeToBuy, unitPrice}) => {
  const totalCost = parseInt((volumeToBuy)) * parseFloat(unitPrice).toPrecision(4)

  return !isNaN(totalCost) && (
    <span className='order-calculation'>
        * {unitPrice} = £ {totalCost.toFixed(2)}
    </span>
  )
}

const isInputTextAllowed = (value, max) =>
  ((value === '') || (value.charAt(0) !== '0' && parseInt(value) <= max && parseInt(value) > 0))

const NumberFormatCustom = ({inputRef, onChange, max, ...other}) => (
  <NumberFormat
    {...other}
    getInputRef={inputRef}
    isAllowed={values => isInputTextAllowed(values.value, max)}
    onValueChange={values => onChange({target: {value: values.value}})}
    allowNegative={false}
    thousandSeparator
    isNumericString
  />
)

export default SelectedOffer

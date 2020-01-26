import React, {useCallback, useEffect, useState} from 'react'
import axios from 'axios'
import produce from 'immer'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import Grid from '@material-ui/core/Grid'
import CssBaseline from '@material-ui/core/CssBaseline'
import Paper from '@material-ui/core/Paper'
import Container from '@material-ui/core/Container'
import StreamedOffers from './components/StreamedOffers'
import {DUX_OFFERS_URL} from './urls'
import SelectedOffer from './components/SelectedOffer'

export const TEST_USER_ID ='RGFya3dpbmcgRHVjawo='

const userBalanceUrl = (userId) => `/api/users/${userId}/balance/`

const fetchData = async (url) => await axios(url)

export const getColorStyle = (color) => {
  switch (color) {
    case 'red':
      return {color: 'red'}
    case 'yellow':
      return {color: 'yellow'}
    default:
      return {color: 'initial'}
  }
}

const App = () => {
  const [offers, setOffers] = useState([])
  const [expiredIds, setExpiredIds] = useState([])
  const [selectedOfferId, setSelectedOfferId] = useState()
  const [balance, setBalance] = useState()
  const [updateBalance, setUpdateBalance] = useState(true)

  // long polling api stream for offers
  const fetchOffers = useCallback(() => {
    fetchData(DUX_OFFERS_URL)
      .then(res => {
        const timeOutId = setTimeout(() => {
          setExpiredIds(produce(draft => {
            draft.push(res.data.id)
          }))
        }, 10000)
        setOffers(produce(draft => {
          draft.push({...res.data, timeOutId})
        }))
      })
      .then(fetchOffers)
  }, [])

  // start long polling when component mounts
  useEffect(() => {
    fetchOffers()
  }, [fetchOffers])

  // remove expired offers
  useEffect(() => {
    setOffers(produce(draft =>
      draft.filter(offer => {
        if (expiredIds.includes(offer.id)) {
          clearTimeout(offer.timeOutId)
          return false
        }
        return true
      })))
  }, [expiredIds])

  // clear order when selected offer expires
  useEffect(() => {
    if (expiredIds.includes(selectedOfferId)) {
      setSelectedOfferId(undefined)
    }
  }, [expiredIds, selectedOfferId])

// vacuum expired ids
  useEffect(() => {
    const id = setInterval(() => setExpiredIds([]), 1000)
    return () => clearInterval(id)
  }, [])

  // update balance after a buy order
  useEffect(() => {
    fetchData(userBalanceUrl(TEST_USER_ID))
      .then(res => {
        setBalance(res.data[`balance`])
        setUpdateBalance(false)
      })
  }, [updateBalance])

  const getSelectedOffer = (id) => produce(offers, draft => {
    const selectedOffer = draft.find(offer => offer.id === id)
    return selectedOffer
  })

  const darkTheme = createMuiTheme({palette: {type: 'dark'}})

  return (
    <MuiThemeProvider theme={darkTheme}>
      <Container className='container'>
        <CssBaseline />
        <header>
          <h1>Dux</h1>
        </header>
        <Grid container spacing={2} wrap='wrap-reverse' className='layout'>
          <Grid item xs={12} sm={12} md={6}>
            <StreamedOffers {...{offers, selectedOfferId, setSelectedOfferId}} />
          </Grid>
          <Grid item xs={12} sm={12} md={6}>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Balance {...{balance}} />
              </Grid>
              <Grid item xs={12}>
                <SelectedOffer {...{selectedOfferId, getSelectedOffer, balance, setUpdateBalance}} />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Container>
    </MuiThemeProvider>
  )
}

const Balance = ({balance}) => (
  <Paper className='active-paper border-blue'>
    <header className='header'>
      <h3>Balance Information</h3>
    </header>
    <Grid container spacing={1} className='balance-details'>
      <Grid item xs={4}>
        User Name:
      </Grid>
      <Grid item xs={8}>
        Test
      </Grid>
      <Grid item xs={4}>
        User ID:
      </Grid>
      <Grid item xs={8}>
        {TEST_USER_ID}
      </Grid>
      <Grid item xs={4}>
        Balance:
      </Grid>
      <Grid item xs={8} className='green'>
        £ {balance}
      </Grid>
    </Grid>
  </Paper>
)

export default App

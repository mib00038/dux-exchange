import React from 'react'
import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
import Divider from '@material-ui/core/Divider'
import capitalize from '@material-ui/core/utils/capitalize'
import Hidden from '@material-ui/core/Hidden'
import Button from '@material-ui/core/Button'
import IconButton from '@material-ui/core/IconButton'
import SelectIcon from '@material-ui/icons/AddCircleOutline'
import SelectedIcon from '@material-ui/icons/Adjust'
import {getColorStyle} from '../App'

const StreamedOffers = ({offers, selectedOfferId, setSelectedOfferId}) => offers && (
  <Paper className='h-100 border-blue active-paper'>
    <header className='header'>
      <h3>Offers Stream</h3>
    </header>
    <ul className='offers-list'>
      <Grid container justify='center' className='list-header'>
        <Grid item xs={3}>Type</Grid>
        <Grid item xs={3}>Price</Grid>
        <Grid item xs={3}>Volume</Grid>
        <Grid item xs={3}/>
      </Grid>
      <Divider color='primary'/>
      {offers.map((offer, index) =>
        <div key={offer.id}>
          <StreamedOffer
            {...offer}
            {...{selectedOfferId, setSelectedOfferId}}
          />
          {offers.length - 1 !== index && <Divider />}
        </div>
      )}
    </ul>
  </Paper>
)

const StreamedOffer = ({id, unitPrice, unitType: {color}, volume, selectedOfferId, setSelectedOfferId}) => {
  const isSelected = (id === selectedOfferId)

  return (
    <li>
      <Grid container alignItems='center' justify='space-between'>
        <Grid item xs={3} style={getColorStyle(color)}>{capitalize(color)}</Grid>
        <Grid item xs={3}>£ {unitPrice}</Grid>
        <Grid item xs={3}>{volume}</Grid>
        <Grid item xs={3}>
          <Hidden xsDown>
            <Button
              className='select-offer-button'
              color={'primary'}
              variant='contained'
              disabled={isSelected}
              onClick={() => setSelectedOfferId(id)}
            >
              {isSelected ? 'Selected' : 'Select'}
            </Button>
          </Hidden>
          <Hidden smUp>
            {!isSelected &&
            <IconButton
              color={'primary'}
              onClick={() => setSelectedOfferId(id)}
            >
              <SelectIcon/>
            </IconButton>
            }
            {isSelected &&
            <IconButton
              color={'primary'}
              disabled={true}
              onClick={() => setSelectedOfferId(id)}
            >
              <SelectedIcon/>
            </IconButton>
            }
          </Hidden>
        </Grid>
      </Grid>
    </li>
  )
}

export default StreamedOffers

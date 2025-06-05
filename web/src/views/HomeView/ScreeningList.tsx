import type { components } from '@/types/openapi/screening';
import { Box, Button, Grid, Paper, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { NavLink } from 'react-router-dom';

type ScreeningListProps = {
  screenings: components['schemas']['ScreeningOutput'][];
};

export const ScreeningList = ({ screenings }: ScreeningListProps) => {
  if (!screenings.length) {
    return <Typography>Brak dostępnych seansów dla filmów w tym kinie.</Typography>;
  }

  return (
    <Grid container spacing={3}>
      {screenings.map((screening) => (
        <Paper key={screening.id} elevation={3} sx={{ p: 2, height: '100%' }}>
          <Typography variant="h6" gutterBottom noWrap>
            {screening.movie?.title}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" noWrap>
            {screening.room?.name} o godzinie {dayjs(screening.startTime).format('HH:mm')}
          </Typography>
          <Box
            component="img"
            src={screening.movie?.photoPath || '/placeholder-movie.png'}
            alt={screening.movie?.title}
            sx={{
              width: '100%',
              height: 180,
              objectFit: 'cover',
              mt: 1,
              borderRadius: 1,
              backgroundColor: '#eee'
            }}
          />
          <NavLink to={`/movie/${screening.movie?.id}`}>
            <Button variant="contained" fullWidth sx={{ mt: 2 }}>
              Zobacz szczegóły
            </Button>
          </NavLink>
        </Paper>
      ))}
    </Grid>
  );
};

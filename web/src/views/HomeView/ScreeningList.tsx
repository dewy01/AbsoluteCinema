import type { components } from '@/types/openapi/screening';
import { mapScreeningsByMovie } from '@/utils/mapScreenings';
import { Box, Button, Paper, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { useMemo } from 'react';
import { NavLink } from 'react-router-dom';

type ScreeningListProps = {
  screenings: components['schemas']['ScreeningOutput'][];
};

export const ScreeningList = ({ screenings }: ScreeningListProps) => {
  if (!screenings.length) {
    return <Typography>Brak dostępnych seansów dla filmów w tym kinie.</Typography>;
  }

  const grouped = useMemo(() => mapScreeningsByMovie(screenings), [screenings]);

  return (
    <Box display="flex" flexDirection="column" gap={2}>
      {grouped.map(({ movie, screenings }) => (
        <Paper
          key={movie?.id}
          elevation={3}
          sx={{
            display: 'flex',
            alignItems: 'flex-start',
            p: 2,
            gap: 2
          }}>
          <Box
            component="img"
            src={movie?.photoPath}
            alt={movie?.title}
            sx={{
              width: 120,
              height: 160,
              objectFit: 'cover',
              borderRadius: 1,
              backgroundColor: '#eee',
              flexShrink: 0
            }}
          />

          <Box sx={{ flex: 1 }}>
            <Typography variant="h6" gutterBottom>
              {movie?.title}
            </Typography>

            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
              {screenings.map((screening) => (
                <NavLink
                  key={screening.id}
                  to={`/screening/${screening.id}`}
                  style={{ textDecoration: 'none' }}>
                  <Button variant="outlined" size="small">
                    {dayjs(screening.startTime).format('HH:mm')}
                  </Button>
                </NavLink>
              ))}
            </Box>
          </Box>
        </Paper>
      ))}
    </Box>
  );
};

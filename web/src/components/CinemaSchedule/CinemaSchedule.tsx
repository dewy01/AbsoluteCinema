import React from 'react';
import {
  Box,
  Button,
  Chip,
  Container,
  Grid,
  MenuItem,
  Select,
  Tab,
  Tabs,
  Typography,
  Card,
  CardMedia,
  CardContent
} from '@mui/material';

const tags = ['DREAM', 'NAPISY', 'DUBBING', 'LEKTOR', 'FILM POLSKI'];
const dates = [
  { label: 'Dzisiaj', date: '5 cze' },
  { label: 'Jutro', date: '6 cze' },
  { label: 'Sob.', date: '7 cze' },
  { label: 'Ndz.', date: '8 cze' },
  { label: 'Pon.', date: '9 cze' },
  { label: 'Wt.', date: '10 cze' },
  { label: 'Śr.', date: '11 cze' }
];

const movies = [
  {
    title: 'Piep*zyć Mickiewicza 2 – Kultura Dostępna',
    age: 'Od 13 lat',
    genre: 'Obyczajowy',
    tag: 'KULTURA DOSTĘPNA',
    time: '18:00',
    image: '/path/to/mickiewicz.jpg'
  },
  {
    title: 'Ballerina. Z uniwersum Johna Wicka',
    age: 'Od 15 lat',
    genre: 'Thriller, akcja',
    tag: 'PRZEDPREMIERA',
    time: '19:00',
    subtitle: 'napisy',
    image: '/path/to/ballerina.jpg'
  }
];

export const CinemaSchedule = () => {
  const [selectedDate, setSelectedDate] = React.useState(0);
  const [genreFilter, setGenreFilter] = React.useState('');

  return (
    <Box sx={{ bgcolor: 'red', color: 'white', py: 3 }}>
      <Container>
        <Typography variant="overline" fontSize={14}>
          REPERTUAR – BIAŁYSTOK HELIOS ALFA
        </Typography>
        <Typography variant="h4" fontWeight="bold">
          TERAZ GRAMY
        </Typography>
        <Typography variant="body2" sx={{ mt: 1 }}>
          Repertuar na kolejny tydzień w <strong>każdy wtorek po godzinie 17:00</strong>
        </Typography>

        {/* Tag filters */}
        <Box sx={{ mt: 2, display: 'flex', gap: 1, flexWrap: 'wrap' }}>
          {tags.map((tag) => (
            <Chip
              key={tag}
              label={tag}
              variant="filled"
              color="default"
              sx={{ bgcolor: 'white', color: 'black' }}
            />
          ))}
        </Box>

        {/* Date selector */}
        <Box sx={{ mt: 3, display: 'flex', gap: 1, flexWrap: 'wrap' }}>
          {dates.map((d, idx) => (
            <Button
              key={idx}
              variant={selectedDate === idx ? 'contained' : 'outlined'}
              sx={{
                bgcolor: selectedDate === idx ? 'white' : 'transparent',
                color: selectedDate === idx ? 'black' : 'white'
              }}
              onClick={() => setSelectedDate(idx)}>
              {d.label}
              <br />
              {d.date}
            </Button>
          ))}
        </Box>

        {/* Genre dropdown */}
        <Box sx={{ mt: 2 }}>
          <Select
            value={genreFilter}
            onChange={(e) => setGenreFilter(e.target.value)}
            displayEmpty
            sx={{ bgcolor: 'white', color: 'black' }}>
            <MenuItem value="">Pokaż wszystkie</MenuItem>
            <MenuItem value="Obyczajowy">Obyczajowy</MenuItem>
            <MenuItem value="Thriller">Thriller</MenuItem>
            <MenuItem value="Akcja">Akcja</MenuItem>
          </Select>
        </Box>
      </Container>

      {/* Film list */}
      <Box sx={{ bgcolor: '#f5f5f5', color: 'black', mt: 3, pt: 2 }}>
        <Container>
          <Tabs value={0} variant="fullWidth" textColor="inherit" indicatorColor="secondary">
            <Tab label="DO POŁUDNIA" />
            <Tab label="PO POŁUDNIU" />
            <Tab label="WIECZOREM" />
          </Tabs>

          <Box sx={{ mt: 2 }}>
            {movies.map((movie, idx) => (
              <Card key={idx} sx={{ display: 'flex', mb: 2 }}>
                <CardMedia
                  component="img"
                  sx={{ width: 100 }}
                  image={movie.image}
                  alt={movie.title}
                />
                <CardContent sx={{ flex: 1 }}>
                  <Typography variant="subtitle1" fontWeight="bold">
                    {movie.title} <Chip label={movie.age} size="small" sx={{ ml: 1 }} />
                  </Typography>
                  <Typography variant="body2" sx={{ mt: 0.5 }}>
                    {movie.genre}{' '}
                    <Chip
                      label={movie.tag}
                      size="small"
                      sx={{
                        ml: 1,
                        bgcolor: movie.tag === 'PRZEDPREMIERA' ? 'red' : 'blue',
                        color: 'white'
                      }}
                    />
                  </Typography>
                </CardContent>
                <Box sx={{ p: 2 }}>
                  <Typography variant="h6" align="center">
                    {movie.time}
                  </Typography>
                  {movie.subtitle && (
                    <Typography variant="caption" display="block" align="center">
                      {movie.subtitle}
                    </Typography>
                  )}
                </Box>
              </Card>
            ))}
          </Box>
        </Container>
      </Box>
    </Box>
  );
};

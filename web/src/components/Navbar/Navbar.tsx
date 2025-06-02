import * as React from 'react';
import {
  AppBar,
  Box,
  Toolbar,
  IconButton,
  Typography,
  Container,
  Avatar,
  Button,
  Tooltip,
  Menu,
  MenuItem,
  Modal
} from '@mui/material';
import { Grid } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import AdbIcon from '@mui/icons-material/Adb';
import LogoIcon from '@/utils/Icons/LogoIcon';
import { useAuth } from '@/contexts';
import { useNavigate } from 'react-router-dom';
import { useCinema } from '@/contexts/CinemaContext';

const cinemaList = ['Cinema City', 'Multikino', 'Helios', 'Test']; // TODO

function Navbar() {
  const { isAuthenticated, logout } = useAuth();
  const navigate = useNavigate();

  const { selectedCinema, isCinemaModalOpen, openCinemaModal, closeCinemaModal, selectCinema } =
    useCinema();

  const [anchorElNav, setAnchorElNav] = React.useState<null | HTMLElement>(null);
  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);

  const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <LogoIcon />
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="/"
            sx={{
              mr: 2,
              display: { xs: 'none', md: 'flex' },
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.3rem',
              color: 'inherit',
              textDecoration: 'none'
            }}>
            ABSOLUTE CINEMA
          </Typography>

          {/* Hamburger (mobile) */}
          <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
            <IconButton size="large" aria-label="menu" onClick={openCinemaModal} color="inherit">
              <MenuIcon />
            </IconButton>
          </Box>

          {/* Ikona i logo (mobile) */}
          <AdbIcon sx={{ display: { xs: 'flex', md: 'none' }, mr: 1 }} />
          <Typography
            variant="h5"
            noWrap
            component="a"
            href="#"
            sx={{
              mr: 2,
              display: { xs: 'flex', md: 'none' },
              flexGrow: 1,
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.3rem',
              color: 'inherit',
              textDecoration: 'none'
            }}>
            LOGO
          </Typography>

          {/* Wybrany kino (desktop) */}
          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' }, alignItems: 'center' }}>
            <Button onClick={openCinemaModal} sx={{ my: 2, color: 'white', display: 'block' }}>
              {selectedCinema ? `Kino: ${selectedCinema}` : 'Wybierz kino'}
            </Button>
          </Box>

          {/* Auth / Avatar */}
          <Box sx={{ flexGrow: 0 }}>
            {isAuthenticated ? (
              <>
                <Tooltip title="Ustawienia">
                  <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                    <Avatar alt="Użytkownik" src="/static/images/avatar/2.jpg" />
                  </IconButton>
                </Tooltip>
                <Menu
                  sx={{ mt: '45px' }}
                  id="menu-appbar"
                  anchorEl={anchorElUser}
                  anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
                  keepMounted
                  transformOrigin={{ vertical: 'top', horizontal: 'right' }}
                  open={Boolean(anchorElUser)}
                  onClose={handleCloseUserMenu}>
                  <MenuItem
                    onClick={() => {
                      handleCloseUserMenu();
                      logout();
                      navigate('/login');
                    }}>
                    <Typography textAlign="center">Wyloguj się</Typography>
                  </MenuItem>
                </Menu>
              </>
            ) : (
              <>
                <Button color="inherit" onClick={() => navigate('/login')}>
                  Zaloguj
                </Button>
                <Button color="inherit" onClick={() => navigate('/register')}>
                  Zarejestruj się
                </Button>
              </>
            )}
          </Box>
        </Toolbar>
      </Container>

      <Modal open={isCinemaModalOpen} onClose={closeCinemaModal}>
        <Box
          sx={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            maxHeight: '70vh',
            overflowY: 'auto',
            width: 600,
            bgcolor: 'background.paper',
            borderRadius: 2,
            boxShadow: 24,
            pt: 0,
            pr: 4,
            pb: 4,
            pl: 4
          }}>
          <Container
            component="div"
            sx={{
              position: 'sticky',
              top: 0,
              zIndex: 1,
              bgcolor: 'background.paper',
              py: 0.5,
              my: 1,
              mb: 1
            }}>
            <Typography variant="h6" sx={{ mb: 1 }}>
              Wybierz kino
            </Typography>
          </Container>

          <Grid container spacing={2}>
            {cinemaList.map((cinema, index) => (
              <Button
                key={index}
                variant="outlined"
                fullWidth
                onClick={() => selectCinema(cinema)}
                sx={{
                  color: 'white',
                  borderColor: 'white',
                  '&:hover': {
                    backgroundColor: 'white',
                    color: 'black',
                    borderColor: 'white'
                  },
                  textTransform: 'none'
                }}>
                {cinema}
              </Button>
            ))}
          </Grid>
        </Box>
      </Modal>
    </AppBar>
  );
}

export default Navbar;

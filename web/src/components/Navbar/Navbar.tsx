import { useUserLogout } from '@/apis/User';
import { selectedCinemaAtom } from '@/atoms/cinemaAtom';
import { useAuth } from '@/contexts';
import LogoIcon from '@/utils/Icons/LogoIcon';
import MenuIcon from '@mui/icons-material/Menu';
import {
  AppBar,
  Avatar,
  Box,
  Button,
  Container,
  Dialog,
  IconButton,
  Menu,
  MenuItem,
  Toolbar,
  Tooltip,
  Typography
} from '@mui/material';
import { useAtom } from 'jotai';
import * as React from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import { CinemaBox } from '../CinemaBox/CinemaBox';

function Navbar() {
  const { isAuthenticated } = useAuth();
  const { mutateAsync } = useUserLogout();
  const navigate = useNavigate();
  const [selectedCinema] = useAtom(selectedCinemaAtom);
  const [isCinemaModalOpen, setCinemaModalOpen] = React.useState(false);
  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);

  const openCinemaModal = () => setCinemaModalOpen(true);
  const closeCinemaModal = () => setCinemaModalOpen(false);

  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <NavLink
            to={'/'}
            style={{ display: 'flex', gap: 4, justifyContent: 'center', alignItems: 'center' }}>
            <LogoIcon />
            <Typography
              variant="h6"
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
          </NavLink>

          {/* Hamburger (mobile) */}
          <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
            <IconButton size="large" aria-label="menu" onClick={openCinemaModal} color="inherit">
              <MenuIcon />
            </IconButton>
          </Box>

          {/* Selected cinema (desktop) */}
          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' }, alignItems: 'center' }}>
            <Button onClick={openCinemaModal} sx={{ my: 2, color: 'white', display: 'block' }}>
              {selectedCinema ? `Kino: ${selectedCinema.name}` : 'Wybierz kino'}
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
                      mutateAsync();
                      handleCloseUserMenu();
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

      <Dialog open={isCinemaModalOpen} onClose={closeCinemaModal}>
        <CinemaBox onClose={closeCinemaModal} />
      </Dialog>
    </AppBar>
  );
}

export default Navbar;

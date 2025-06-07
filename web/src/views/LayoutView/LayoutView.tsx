import { Footer } from '@/components/Footer/Footer';
import Navbar from '@/components/Navbar/Navbar';
import Progress from '@/components/Progress/Progress';
import { Box } from '@mui/material';
import { Suspense } from 'react';
import { Outlet } from 'react-router-dom';

export const LayoutView = () => {
  return (
    <Box
      component="section"
      sx={{
        display: 'flex',
        flexDirection: 'column',
        maxHeight: '100vh',
        minHeight: '100vh'
      }}>
      <Navbar />

      <Box
        component="main"
        sx={{
          display: 'flex',
          flexDirection: 'column',
          flex: 1,
          flexGrow: 1,
          boxSizing: 'border-box',
          justifyContent: 'center',
          alignItems: 'center'
        }}>
        <Suspense fallback={<Progress />}>
          <Outlet />
        </Suspense>
      </Box>
      <Footer />
    </Box>
  );
};

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
        height: '100vh',
        width: '100vw',
        overflow: 'hidden'
      }}>
      <Box component="header" sx={{ flexShrink: 0 }}>
        <Navbar />
      </Box>

      <Box
        component="main"
        sx={{
          flexGrow: 1,
          overflow: 'auto',
          p: 2
        }}>
        <Suspense fallback={<Progress />}>
          <Outlet />
        </Suspense>
      </Box>
    </Box>
  );
};

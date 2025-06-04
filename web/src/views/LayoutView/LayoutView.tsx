import Navbar from '@/components/Navbar/Navbar';
import Progress from '@/components/Progress/Progress';
import { Outlet } from 'react-router-dom';
import { Box } from '@mui/material';
import { Suspense } from 'react';

export const LayoutView = () => {
  return (
    <Box
      component="section"
      sx={{
        display: 'flex',
        maxHeight: '100vh',
        minHeight: '100vh',
        minWidth: '100vw',
        maxWidth: '100vw',
        flexDirection: 'column'
      }}>
      <Navbar />
      <Box sx={{ display: 'flex', flexGrow: 1, flex: 1 }}>
        <Suspense fallback={<Progress />}>
          <Outlet />
        </Suspense>
      </Box>
    </Box>
  );
};

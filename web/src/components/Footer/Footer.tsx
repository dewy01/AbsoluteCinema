import { Box, Link, Paper, Typography } from '@mui/material';
import { useState } from 'react';
import { DeleteDialog } from './DeleteDialog';
import { UpdateDialog } from './UpdateDialog';

export const Footer = () => {
  const [deleteOpen, setDeleteOpen] = useState(false);
  const [updateOpen, setUpdateOpen] = useState(false);

  return (
    <>
      <Paper
        component="footer"
        sx={{
          p: 2,
          mt: 'auto',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          gap: 1
        }}>
        <Typography variant="subtitle1" fontWeight="bold">
          Rezerwacje
        </Typography>
        <Box display="flex" gap={2}>
          <Link
            component="button"
            variant="body2"
            onClick={() => setDeleteOpen(true)}
            sx={{ cursor: 'pointer' }}>
            Anuluj rezerwacje
          </Link>
          <Link
            component="button"
            variant="body2"
            onClick={() => setUpdateOpen(true)}
            sx={{ cursor: 'pointer' }}>
            Edytuj rezerwacje
          </Link>
        </Box>
      </Paper>

      <DeleteDialog open={deleteOpen} onClose={() => setDeleteOpen(false)} />
      <UpdateDialog open={updateOpen} onClose={() => setUpdateOpen(false)} />
    </>
  );
};

import { useUserLogin } from '@/apis/User';
import { userLoginSchema, type UserLogin } from '@/schemas/UserLogin';
import { zodResolver } from '@hookform/resolvers/zod';
import { Box, Button, Paper, TextField, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';

export const LoginView = () => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset
  } = useForm<UserLogin>({
    resolver: zodResolver(userLoginSchema),
    mode: 'onChange',
    defaultValues: {
      email: '',
      password: ''
    }
  });

  const { mutateAsync } = useUserLogin();

  const onSubmit = async (data: UserLogin) => {
    mutateAsync(data).then(() => {
      reset();
    });
  };

  return (
    <Box
      sx={{
        display: 'flex',
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%',
        height: '100%'
      }}>
      <Paper
        elevation={6}
        sx={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          minWidth: '350px',
          minHeight: '300px',
          p: 4,
          px: 6,
          gap: 4,
          borderRadius: 3
        }}>
        <Typography variant="h5" component="h1" align="center" fontWeight={700}>
          Zaloguj się
        </Typography>

        <Box
          component="form"
          onSubmit={handleSubmit(onSubmit)}
          sx={{
            display: 'flex',
            flex: 1,
            flexDirection: 'column',
            alignItems: 'center',
            gap: 2
          }}
          noValidate
          autoComplete="off">
          <TextField
            label="Adres email"
            {...register('email')}
            error={!!errors.email}
            helperText={errors.email?.message}
            sx={{ width: '250px' }}
          />

          <TextField
            label="Hasło"
            type="password"
            {...register('password')}
            error={!!errors.password}
            helperText={errors.password?.message}
            sx={{ width: '250px' }}
          />

          <Button type="submit" variant="contained" color="primary" disabled={isSubmitting}>
            Zaloguj się
          </Button>
        </Box>

        {/* <NavLink to="/" style={{ fontSize: '0.875rem' }}>
          Nie pamiętasz hasła?
        </NavLink> */}
      </Paper>
    </Box>
  );
};

import { callUserRegister } from '@/apis/User';
import { userRegisterSchema, type UserRegister } from '@/schemas/UserRegister';
import { zodResolver } from '@hookform/resolvers/zod';
import { Box, Button, Paper, TextField, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import { NavLink } from 'react-router-dom';

export const RegisterView = () => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset
  } = useForm<UserRegister>({
    resolver: zodResolver(userRegisterSchema),
    mode: 'onChange',
    defaultValues: {
      name: '',
      email: '',
      password: '',
      confirmPassword: ''
    }
  });

  const { mutateAsync } = callUserRegister();

  const onSubmit = async (data: UserRegister) => {
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
          minWidth: '350px',
          minHeight: '450px',
          p: 4,
          px: 6,
          gap: 4,
          borderRadius: 3
        }}>
        <Typography variant="h5" component="h1" align="center" fontWeight={700}>
          Zarejestruj się
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
            label="Imię"
            {...register('name')}
            error={!!errors.name}
            helperText={errors.name?.message}
            sx={{ width: '250px' }}
          />

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

          <TextField
            label="Powtórz hasło"
            type="password"
            {...register('confirmPassword')}
            error={!!errors.confirmPassword}
            helperText={errors.confirmPassword?.message}
            sx={{ width: '250px' }}
          />

          <Button type="submit" variant="contained" color="primary" disabled={isSubmitting}>
            Zarejestruj się
          </Button>
        </Box>

        <NavLink to="/login" style={{ fontSize: '0.875rem' }}>
          Masz już konto? Zaloguj się
        </NavLink>
      </Paper>
    </Box>
  );
};

import { Box, TextField, Button, Paper, Typography, Link } from '@mui/material';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { userRegisterSchema, type UserRegister } from '@/schemas/UserRegister';

export const RegisterView = () => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting }
  } = useForm<UserRegister>({
    resolver: zodResolver(userRegisterSchema),
    mode: 'onChange'
  });

  const onSubmit = (data: UserRegister) => {
    console.log('Dane z formularza rejestracji:', data);
  };

  return (
    <Box
      component="section"
      sx={{
        display: 'flex',
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%'
      }}>
      <Paper
        elevation={6}
        sx={{
          p: 4,
          borderRadius: 3,
          backgroundColor: 'background.paper',
          minWidth: '300px'
        }}>
        <Typography variant="h5" component="h1" align="center" sx={{ mb: 2, fontWeight: 'bold' }}>
          Zarejestruj się
        </Typography>

        <Box
          component="form"
          onSubmit={handleSubmit(onSubmit)}
          sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            '& .MuiTextField-root': { m: 1, width: '25ch' }
          }}
          noValidate
          autoComplete="off">
          <TextField
            label="Adres email"
            required
            {...register('email')}
            error={!!errors.email}
            helperText={errors.email?.message}
          />

          <TextField
            label="Hasło"
            type="password"
            required
            {...register('password')}
            error={!!errors.password}
            helperText={errors.password?.message}
          />

          <TextField
            label="Powtórz hasło"
            type="password"
            required
            {...register('confirmPassword')}
            error={!!errors.confirmPassword}
            helperText={errors.confirmPassword?.message}
          />

          <Button
            type="submit"
            variant="contained"
            color="primary"
            sx={{ mt: 2, width: '25ch' }}
            disabled={isSubmitting}>
            Zarejestruj się
          </Button>

          <Link href="/login" underline="hover" sx={{ mt: 1, fontSize: '0.875rem' }}>
            Masz już konto? Zaloguj się
          </Link>
        </Box>
      </Paper>
    </Box>
  );
};

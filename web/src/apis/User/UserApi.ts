export interface User {
  firstName: string;
  email: string;
  image?: string;
  role: 'user' | 'admin';
}

export interface Accesstoken {
  accessToken: string;
  refreshToken: string;
}

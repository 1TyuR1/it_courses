import { Storage, STORAGE_KEYS } from '../utils/storage';

const API_BASE_URL = 'https://your-api-url.com/api'; 

export interface AuthResponse {
  token: string;
  user: {
    id: string;
    email: string;
  };
}

export interface LoginData {
  email: string;
  password: string;
}

export interface RegisterData {
  email: string;
  password: string;
  confirmPassword: string;
}

class AuthService {
  private token: string | null = null;

  constructor() {
    this.loadToken();
  }

  private async loadToken() {
    this.token = await Storage.getItem(STORAGE_KEYS.AUTH_TOKEN);
  }

  async register(userData: RegisterData): Promise<AuthResponse> {
    if (userData.password !== userData.confirmPassword) {
      throw new Error('Пароли не совпадают');
    }

    if (userData.password.length < 6) {
      throw new Error('Пароль должен содержать минимум 6 символов');
    }

    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: userData.email,
        password: userData.password,
      }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.message || 'Ошибка регистрации');
    }

    const data: AuthResponse = await response.json();
    await this.setToken(data.token);
    return data;
  }

  async login(credentials: LoginData): Promise<AuthResponse> {
    const response = await fetch(`${API_BASE_URL}/auth/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.message || 'Ошибка входа');
    }

    const data: AuthResponse = await response.json();
    await this.setToken(data.token);
    return data;
  }

  async logout(): Promise<void> {
    this.token = null;
    await Storage.removeItem(STORAGE_KEYS.AUTH_TOKEN);
    await Storage.removeItem(STORAGE_KEYS.USER_DATA);
  }

  async setToken(token: string): Promise<void> {
    this.token = token;
    await Storage.setItem(STORAGE_KEYS.AUTH_TOKEN, token);
  }

  getToken(): string | null {
    return this.token;
  }

  isAuthenticated(): boolean {
    return !!this.token;
  }

  async authenticatedFetch(url: string, options: RequestInit = {}) {
    if (!this.token) {
      throw new Error('Not authenticated');
    }

    const headers = {
      'Authorization': `Bearer ${this.token}`,
      'Content-Type': 'application/json',
      ...options.headers,
    };

    return fetch(url, { ...options, headers });
  }
}

export const authService = new AuthService();
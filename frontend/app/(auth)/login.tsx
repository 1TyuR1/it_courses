import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { useAuth } from '../../hooks/auth-context';
import { AuthForm } from '../../components/AuthForm';
import { router } from 'expo-router';

export default function LoginScreen() {
  const { login, isLoading } = useAuth();

  return (
    <View style={{ flex: 1 }}>
      <AuthForm
        type="login"
        onSubmit={login}
        isLoading={isLoading}
      />
      
      <View style={{ padding: 20, alignItems: 'center' }}>
        <Text style={{ marginBottom: 10 }}>Нет аккаунта?</Text>
        <TouchableOpacity onPress={() => router.push('../(auth)/register')}>
          <Text style={{ color: '#007AFF', fontWeight: '600' }}>
            Зарегистрироваться
          </Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}
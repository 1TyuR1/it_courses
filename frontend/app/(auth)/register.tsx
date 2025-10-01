import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { useAuth } from '../../hooks/auth-context';
import { AuthForm } from '../../components/AuthForm';
import { router } from 'expo-router';

export default function RegisterScreen() {
  const { register, isLoading } = useAuth();

  return (
    <View style={{ flex: 1 }}>
      <AuthForm
        type="register"
        onSubmit={register}
        isLoading={isLoading}
      />
      
      <View style={{ padding: 20, alignItems: 'center' }}>
        <Text style={{ marginBottom: 10 }}>Уже есть аккаунт?</Text>
        <TouchableOpacity onPress={() => router.push('../(auth)/login')}>
          <Text style={{ color: '#007AFF', fontWeight: '600' }}>
            Войти
          </Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}
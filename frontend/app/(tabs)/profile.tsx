import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { useAuth } from '../../hooks/auth-context';

export default function ProfileScreen() {
  const { user, logout } = useAuth();

  return (
    <View style={{ flex: 1, padding: 20 }}>
      <Text style={{ fontSize: 24, fontWeight: 'bold', marginBottom: 20 }}>
        Профиль
      </Text>
      
      <Text style={{ fontSize: 16, marginBottom: 10 }}>
        Email: {user?.email}
      </Text>

      <TouchableOpacity
        style={{
          backgroundColor: '#FF3B30',
          padding: 16,
          borderRadius: 8,
          alignItems: 'center',
          marginTop: 20,
        }}
        onPress={logout}
      >
        <Text style={{ color: '#FFFFFF', fontSize: 16, fontWeight: '600' }}>
          Выйти
        </Text>
      </TouchableOpacity>
    </View>
  );
}
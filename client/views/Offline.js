import React from 'react';
import { StyleSheet, Text, View, Image } from 'react-native';

export default (props) => {
  console.log(props);
  return (
    <View>
      <Image
        source={require('../images/offline.png')}
        style={{ marginTop: 20, width: '100%', height: '75%' }}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
});

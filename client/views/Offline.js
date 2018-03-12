import React from 'react';
import { StyleSheet, View, Image } from 'react-native';

export default (props) => {
  console.log(props);
  return (
    <View>
      <Image source={require('../images/offline.png')} style={styles.image} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  image: {
    marginTop: 20,
    width: '100%',
    height: '75%',
  },
});

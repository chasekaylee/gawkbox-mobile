import React, { Component } from 'react';
import { StyleSheet, Text, Button } from 'react-native';
import { Card } from 'react-native-elements';

export default (props) => {
  const { item, navigation: { navigate } } = props;
  return (
    <Card title={item.channel.name} image={{ uri: item.preview.large }}>
      <Text style={styles.status}>{item.channel.status}</Text>
      <Button
        icon={{ name: 'code' }}
        backgroundColor="#03A9F4"
        fontFamily="Lato"
        buttonStyle={styles.button}
        title="Watch Stream"
        onPress={() => navigate('StreamDetails', { stream: item })}
      />
    </Card>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
  status: {
    marginBottom: 10,
  },
  button: {
    borderRadius: 0,
    marginLeft: 0,
    marginRight: 0,
    marginBottom: 0,
  },
});

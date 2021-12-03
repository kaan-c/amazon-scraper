import logo from './logo.svg';
import './App.css';
import React from 'react';

class pageComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {address: "", dataJson: ""};
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this); 
  }
  handleSubmit(event) {
    console.log("address is " + this.state.address);
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ address: this.state.address })
    };
    fetch("http://localhost:10000/address", requestOptions)
      .then((response) => response.json())
      .then((data) => this.setState({address:"", dataJson: data[0].title}));
    event.preventDefault();
  }

  handleChange(event) {
    this.setState({address: event.target.value});  
  }

  render() {
    return (
    <html>
      <form onSubmit={this.handleSubmit}>
            <label>Address:</label>
            <input type="text" id="url" value={this.state.address} onChange={this.handleChange} />
            <br />
            <button type="submit">Submit</button>
            <br/>
            <label>{this.state.dataJson}</label>
      </form>
    </html>);
  }
}

export default pageComponent;

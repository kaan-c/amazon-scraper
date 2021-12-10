import logo from './logo.svg';
import './App.css';
import React from 'react';

class pageComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {address: "", productList: []};
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this); 
  }
  handleApiOutput(data) {
    let productList = [];
    data.map(dataRow => {
      productList.push({"title": dataRow.title, "price": dataRow.price});
    });
    this.setState({"address": this.state.address, "productList": productList});
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
      .then((data) => this.handleApiOutput(data));
    event.preventDefault();
  }
  handleChange(event) {
    this.setState({address: event.target.value});  
  }

  render() {
    return (
    <div>
      <form onSubmit={this.handleSubmit}>
        <label>Address:</label>
        <input type="text" id="url" value={this.state.address} onChange={this.handleChange} />
        <br />
        <button type="submit">Submit</button>
        <br/>
      </form>
      <div>
        {this.state.productList.map((item, index) => (
          <p><input type="checkbox" /><label key={index}>{item.title} {item.price}</label></p>
        ))}
      </div>
    </div>);
  }
}

export default pageComponent;

import './App.css';
import React, { } from 'react';

class pageComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {address: "", productList: []};
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleSelectProduct = this.handleProductSelect.bind(this);
  }
  handleApiOutput(data) {
    let productList = [];
    data.map(dataRow => {
      productList.push({"title": dataRow.title, "price": dataRow.price, "selected": false});
      return productList;
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
  handleProductSelect(event, index) {
    const updatedProduct = {"title": this.state.productList[index].title, 
                            "price": this.state.productList[index].price, 
                            "selected": event.target.checked};
    this.setState(prevState => {
      prevState.productList[index] = updatedProduct;
      return { prevState };
    });
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
          <p><input type="checkbox" onChange={(event) => this.handleProductSelect(event, index)} /><label key={index}>{item.title} {item.price}</label></p>
        ))}
      </div>
    </div>);
  }
}

export default pageComponent;

import { Component } from "react";

import { Navigate } from "react-router-dom";
import AuthService from "../../../services/auth.service";
import IUser from "../../../types/user.type";
import axios from "axios";


type Props = {};

type State = {
  context: null
}

export default class Search extends Component<Props, State> {
  constructor(props: Props) {
    super(props);

    this.searchContext = this.searchContext.bind(this)

    this.state = {
      context: null
    };
  }

  async searchContext(){
    const headers = { 
        "Authorization": `Bearer ${AuthService.getCurrentUser().token}`,
    };

    const body = {
        "search": "word"
    };

    await axios.post("http://localhost:8080/api/v1/search", { body },{ headers })
    .then(res => {
        let context = res.data.context
        this.setState({ context });
    })
  }

  render() {
    return (
      <div className="container">
        <p>
            <strong>Email:</strong>{" "}
            { AuthService.getCurrentUser().user.email }
        </p>
        
        <input type="text" onClick={ this.searchContext }
      </div>
    );
  }
}
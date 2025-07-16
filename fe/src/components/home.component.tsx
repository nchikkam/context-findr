import { Component } from "react";
import AuthService from "../services/auth.service";


type Props = {};

type State = {
  content: string;
}

export default class Home extends Component<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      content: "Simple Context Finder App"
    };
  }

  
  render() {
    return (
      <div className="container">
        <header className="jumbotron">
          
          { AuthService.getCurrentUser() ? "User Logged in" : "User not Logged in" }
          
          <h3>{this.state.content}</h3>
        </header>
      </div>
    );
  }
}
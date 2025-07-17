import { Component, KeyboardEvent } from "react";

import AuthService from "../../../services/auth.service";
import axios from "axios";

import ReactJson, { ReactJsonViewProps } from "react-json-view";


type Props = {};

type State = {
  matches: object | null
}

export default class Search extends Component<Props, State> {
  constructor(props: Props) {
    super(props);

    this.searchContext = this.searchContext.bind(this)
    this.handleEnter = this.handleEnter.bind(this)

    this.state = {
      matches: null
    };
  }

  searchContext(query: string){
    const headers = { 
        "Authorization": `Bearer ${AuthService.getCurrentUser().token}`,
    };

    const url = "http://localhost:8080/api/v1/search";
    
    axios.get(`${url}?q=${query}`, { headers })
    .then(res => {
        let matches = res.data.matches
        this.setState({ matches });
    })
  }

  //handleEnter(e: ChangeEvent<HTMLInputElement>): void {
  handleEnter(e: KeyboardEvent<HTMLInputElement>): void {
    const target = (e.target as HTMLInputElement)
    
    if (e.key === "Enter") {
        this.searchContext(target.value);
    }
  }

  render() {
    return (
      <div className="container">
        <p>
            <strong>Email:</strong>{" "}
            { AuthService.getCurrentUser().user.email }

        </p>
        
        <input onKeyDown={this.handleEnter} />

        <br/><br/>
        <p>
          Snippets and Context Results:
        </p>

        <hr />
          { 
            this.state.matches ?
              <ReactJson 
                src={(this.state.matches as ReactJsonViewProps)} 
                theme={"bright:inverted"}
                iconStyle="square"
                displayDataTypes={false}
                displayObjectSize={true}
              />
              : 
              <p>Context Not Found, Please try some random string</p>
          }
      </div>
    );
  }
}
import { Component } from "react";
import { Navigate } from "react-router-dom";
import AuthService from "../services/auth.service";
import IUser from "../types/user.type";
import Upload from "./lib/uploader/upload.component";
import FileUploads from "./lib/uploader/fileuploads.component";


type Props = {};

type State = {
  redirect: string | null,
  userReady: boolean,
  currentUser: IUser & { token: string }
}
export default class Profile extends Component<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      redirect: null,
      userReady: false,
      currentUser: { token: "" }
    };
  }

  componentDidMount() {
    const currentUser = AuthService.getCurrentUser();

    if (!currentUser) this.setState({ redirect: "/home" });
    this.setState({ currentUser: currentUser, userReady: true })
  }

  render() {
    if (this.state.redirect) {
      return <Navigate to={this.state.redirect} />
    }

    return (
      <div className="container">
        {(this.state.userReady) ?
          <div>

            <p>
              Registered User.
            </p>
            <p>
              <strong>Email:</strong>{" "}
              { AuthService.getCurrentUser().user.email }
            </p>

            <Upload />
            <FileUploads />

          </div> : null}
      </div>
    );
  }
}
import axios from "axios";
import { Component } from "react";

import AuthService from "../../../services/auth.service";
import { Server } from "../../../common/helper";


type Props = {};

type State = {
  files: []
}

interface File {
    _id: string
    name: string
    email: string
}

export default class FileUploads extends Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            files: []
        };
    }

    async componentDidMount() {
        const headers = { 
            "Authorization": `Bearer ${AuthService.getCurrentUser().token}`,
        };

        await axios.get(`${Server.baseURL}/api/v1/uploads`, { headers })
        .then(res => {
            let files = res.data.files
            this.setState({ files });
        })
    }

    render() {
        if (this.state.files){
            return (
            <table>
                <tbody>
                    <tr>
                    <th>File</th>
                    </tr>
                        {
                            this.state.files
                            .map((file: File, idx: number) => {
                                return <tr><td key={idx}>{file.name}</td></tr>
                            })
                        }
                </tbody>
            </table>
        )
        } else {
            return <div></div>
        }
        
    }
}
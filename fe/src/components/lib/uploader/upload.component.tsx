import axios from "axios";
import { Component } from "react";

import AuthService from "../../../services/auth.service";
import { Server } from "../../../common/helper";

export default class Upload extends Component {
    constructor(props: any){
        super(props);
        this.uploadFile = this.uploadFile.bind(this);
    }
    
    uploadFile(event: React.ChangeEvent<HTMLInputElement>){
        const selectedFiles = event.target.files as FileList;
        let formData = new FormData();
        formData.append("file", selectedFiles?.[0]);

        const headers = { 
            "Content-Type": "multipart/form-data",
            "Authorization": `Bearer ${AuthService.getCurrentUser().token}`,
        };

        axios.post(`${Server.baseURL}/api/v1/upload`, formData, { headers });
    }

    render (){
        return (
            <div>
                <input type="file" onChange={ this.uploadFile } />
            </div>
        )
    }
}
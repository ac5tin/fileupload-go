import React from 'react';
import Upload from '../upload';


const Landing = ()=>(
    <React.Fragment>
         <div className="section header">
            <h2 className="title">File Uploader</h2>
        </div>

        <div className="section">
            <h3 className="subtitle">Simple file sharing</h3>
            <Upload />
        </div>
    </React.Fragment>
    
)

export default Landing;

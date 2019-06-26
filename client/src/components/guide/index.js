import React from 'react';


const Guide  = ()=>(
    <React.Fragment>
        

        <div className="row">
            <p className="full_width">Download</p>
            <pre className="code">
                <code>$ curl {process.env.REACT_APP_ADDR}/api/file/d/FqigXPc8wD -o foo.txt</code>
            </pre>
        </div>
        

        <div className="row">
            <p className="full_width">Upload</p>
            <pre className="code">
                <code>$ curl -X PUT {process.env.REACT_APP_ADDR}/api/file/u -F file=@foo.txt</code>
            </pre>
        </div>
        

    </React.Fragment>
    
)

export default Guide;
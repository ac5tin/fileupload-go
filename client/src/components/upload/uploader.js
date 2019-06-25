import React from 'react';
import Dropzone from 'react-dropzone';
import { formatBytes } from 'usefuljs';


const Uploader = ({ maxSize=1073741824 , minSize=1, onDrop=()=>{} , onReject=()=>{} })=>{
    const _onDrop = React.useCallback(acceptedFiles => {acceptedFiles.length && onDrop(acceptedFiles[0])},[onDrop]);

    const _onReject = files => onReject(files);

    return (
        <span>
            <Dropzone
                minSize={minSize}
                maxSize={maxSize}
                onDrop={_onDrop}
                onDropRejected={_onReject}
            >
                {({getRootProps, getInputProps,isDragActive}) => (
                    <section>
                        <div {...getRootProps()} className={"border p-1 filedropper" + (isDragActive ? " dragging" :"")}>
                            <input {...getInputProps()} />
                            <React.Fragment>
                            {
                                isDragActive ? 
                                <p>Drop your file here</p>: 
                                <p>Drag a file here (max size is {formatBytes(maxSize)}) </p>
                            }
                            </React.Fragment>
                        </div>
                    </section>
                )}
            </Dropzone>
        </span>
    )
}
export default Uploader;
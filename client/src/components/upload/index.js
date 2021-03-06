import React from 'react';
import { toast } from 'react-toastify';
import { formatBytes } from 'usefuljs';
import Uploader from './uploader';

const Upload = ()=>{
    const [files,setFiles] = React.useState([]); // file is uploaded/done
    const [uploading,setUploading] = React.useState(false);//uploading file

    const onDrop = async f =>{
        await setUploading(true); //uploading file
        const t = await toast('Uploading file',{ autoClose: false , type: toast.TYPE.WARNING });
        // instantiate ws connection
        const ws = new WebSocket(process.env.REACT_APP_WS_ENDPOINT);
        ws.binaryType = "arraybuffer";

        let reader = new FileReader();
        reader.onload = async e =>{
            await ws.send(new Uint8Array(e.target.result));
        }
        // send start signal
        // file API -> read as arraybuffer
        // send arraybuffer
        // send end signal
        // wait for ws done signal with download link
        ws.onopen = async()=>{
            await ws.send(JSON.stringify({filename: f.name, size: f.size}))
            reader.readAsArrayBuffer(f);
        }


        ws.onclose = ()=>{
            //console.log('upload training data ws closed');//debug
            reader = null;//gc
        }
        
        ws.onmessage = msg =>{
            const { result , id } = JSON.parse(msg.data);
            if(result === "success" && id){
                toast.update(t,{render:"Successfully uploaded file" ,type: toast.TYPE.SUCCESS , autoClose: 3000 });
                f.dlink = `${process.env.REACT_APP_ADDR}/api/file/d/${id}`;
                setFiles([...files,f]);
            }else{
                toast.update(t,{render:"Failed to upload file" ,type: toast.TYPE.SUCCESS , autoClose: 3000 });
            }
            setUploading(false);// done
            ws.close();
        }

        
    }

    const onReject = () =>{
        toast.error('Failed to upload file');
        setUploading(false);
    }

    return (
        <React.Fragment>
            
            {
                uploading ? null :
                <div className="row">
                    <Uploader onDrop={onDrop} onReject={onReject} maxSize={parseInt(process.env.REACT_APP_MAX_SIZE)} />
                </div>
                
            }
            {
                files.length ?
                    <table>
                        <thead>
                            <tr>
                                <th>File</th>
                                <th>Status</th>
                                <th>Link</th>
                                <th>Size</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                files.map(f => 
                                    <tr key={f.name+String(f.size)}>
                                        <td>{f.name}</td>
                                        <td>{uploading ? "Uploading ..." : "Done"}</td>
                                        <td>{f.dlink ? <a href={f.dlink}  rel="noopener noreferrer" target="_blank">{f.dlink}</a> : ""}</td>
                                        <td>{formatBytes(f.size)}</td>
                                    </tr>
                                )
                            }
                        </tbody>
                    </table>

                :
                    null
            }
                
            
        </React.Fragment>
            
    )
}
export default Upload;

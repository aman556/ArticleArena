import {GoogleLogin, GoogleOAuthProvider} from '@react-oauth/google';
import jwt_decode from 'jwt-decode';
import {useNavigate } from 'react-router-dom';

const GoogleSignUp: React.FC = () => {
    const navigate = useNavigate();
    return (
        <div>
            <GoogleOAuthProvider clientId='1053340939906-a88i2gbtul1jgahmgrj24ou8jlje98hi.apps.googleusercontent.com'>
                <GoogleLogin
                    onSuccess={(credentialResponse: any) => {
                        var decoded:any = jwt_decode(credentialResponse.credential);
                        console.log(decoded.name);
                        if(decoded.email_verified){
                            navigate('/profile')
                        }
                    }}
                    theme="filled_blue"
                    shape="circle"
                    onError={() => {
                        console.log('Login Failed');
                    }}
                    useOneTap
                />
            </GoogleOAuthProvider>
        </div>
    )
}

export default GoogleSignUp;
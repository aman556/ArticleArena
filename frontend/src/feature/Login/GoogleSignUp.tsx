import {GoogleLogin, GoogleOAuthProvider} from '@react-oauth/google';
import jwt_decode from 'jwt-decode';

const GoogleSignUp: React.FC = () => {
    return (
        <div>
            <GoogleOAuthProvider clientId='1053340939906-a88i2gbtul1jgahmgrj24ou8jlje98hi.apps.googleusercontent.com'>
                <GoogleLogin
                    onSuccess={(credentialResponse: any) => {
                        var decoded:any = jwt_decode(credentialResponse.credential);
                        console.log(decoded.name);
                    }}
                    onError={() => {
                        console.log('Login Failed');
                    }}
                />
            </GoogleOAuthProvider>;
        </div>
    )
}

export default GoogleSignUp;
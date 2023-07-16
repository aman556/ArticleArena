import useStyles from './styles';
import { googleLogout } from '@react-oauth/google';

const LogoutButton: React.FC = () => {
    const styles = useStyles();
    
    return (
        <div className={styles.logut}>
            <button type="button" onClickCapture={Logout}>LogOut</button>
        </div>
    )
}

function Logout() {
    
}

export default LogoutButton;
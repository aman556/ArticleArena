import useStyles from './styles';
import { Button } from '../../styles/Button';
import { NavLink } from 'react-router-dom';

const LogoutButton: React.FC = () => {
    const styles = useStyles();
    
    return (
        <div className={styles.logout}>
            <NavLink to="/">
                <button type="button" onClickCapture={Logout}>LogOut</button>
            </NavLink>
        </div>
    )
}

function Logout() {
    
}

export default LogoutButton;
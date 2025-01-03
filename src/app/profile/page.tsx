'use client';

import { Profile } from '@/types/User';
import { createClient } from '@/utils/supabase/server';
import { useEffect, useState } from 'react';

export default function UserProfile() {
    
    const [profile, setProfile] = useState<Profile | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    
    useEffect(() => {
        const fetchProfile = async() => {
            try {
                const supabase = await createClient();
                const {data: { user }} = await supabase.auth.getUser();

                if(user) {
                    const {data: profileData, error} = await supabase
                        .from('profiles')
                        .select('*')
                        .eq('id', user.id)
                        .single();
                    if(error) {
                        console.error("Error fetching profile: ", error);
                    }
                    setProfile(profileData);
                }
            } catch(error) {
                console.error("Error fetching user:", error);
            } finally {
                setLoading(false);
            }
        };

        fetchProfile();
    }, []);

    if(!profile) {
        return (
            <>
                <div>
                    Loading...
                </div>
            </>
        );
    }

    return (
        <>
            <h1>Hi! Welcome {profile.name}</h1>
        </>
    );
}